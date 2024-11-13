#!/usr/bin/env python3

import re
import sys
import argparse
from pathlib import Path

class CommentStripper:
    PATTERNS = {
        '.py': {
            'single': '#',
            'multi': ['"""', "'''"]
        },
        '.sh': {
            'single': '#',
            'multi': None
        },
        '.rb': {
            'single': '#',
            'multi': ['=begin', '=end']
        },
        '.c': {
            'single': '//',
            'multi': ['/*', '*/']
        },
        '.cpp': {
            'single': '//',
            'multi': ['/*', '*/']
        },
        '.js': {
            'single': '//',
            'multi': ['/*', '*/']
        },
        '.java': {
            'single': '//',
            'multi': ['/*', '*/']
        },
        '.html': {
            'single': None,
            'multi': ['<!--', '-->']
        },
        '.css': {
            'single': None,
            'multi': ['/*', '*/']
        },
        '.sql': {
            'single': '--',
            'multi': ['/*', '*/']
        },
        '.lisp': {
            'single': ';',
            'multi': None
        },
        '.lua': {
            'single': '--',
            'multi': ['--[[', ']]']
        },
        '.go': {
            'single': '//',
            'multi': ['/*', '*/']
        },
        '.rs': {
            'single': '//',
            'multi': ['/*', '*/']
        }
    }

    def __init__(self, keep_docstrings=False):
        self.keep_docstrings = keep_docstrings

    def strip_comments(self, content, file_ext):
        if file_ext not in self.PATTERNS:
            return content

        pattern = self.PATTERNS[file_ext]
        result = []
        in_multiline = False
        multiline_start = None

        lines = content.split('\n')
        i = 0
        while i < len(lines):
            line = lines[i]

            if pattern['multi']:
                if not in_multiline:
                    if pattern['multi'][0] in line:
                        if self.keep_docstrings and file_ext == '.py' and i > 0:
                            prev_line = lines[i-1].strip()
                            if prev_line.startswith('def ') or prev_line.startswith('class '):
                                result.append(line)
                                i += 1
                                continue

                        in_multiline = True
                        multiline_start = pattern['multi'][0]
                        before_comment = line.split(pattern['multi'][0])[0]
                        if before_comment.strip():
                            result.append(before_comment)
                    else:
                        if pattern['single']:
                            if pattern['single'] in line:
                                code = line.split(pattern['single'])[0]
                                if code.strip():
                                    result.append(code)
                            else:
                                result.append(line)
                        else:
                            result.append(line)
                else:
                    if pattern['multi'][1] in line:
                        in_multiline = False
                        after_comment = line.split(pattern['multi'][1])[1]
                        if after_comment.strip():
                            result.append(after_comment)
            else:
                if pattern['single'] and pattern['single'] in line:
                    code = line.split(pattern['single'])[0]
                    if code.strip():
                        result.append(code)
                else:
                    result.append(line)

            i += 1

        return '\n'.join(result)

def main():
    parser = argparse.ArgumentParser(description='Strip comments from source code files.')
    parser.add_argument('file', help='File to process')
    parser.add_argument('--keep-docstrings', action='store_true',
                      help='Keep Python docstrings for classes and functions')
    parser.add_argument('--output', '-o', help='Output file (default: stdout)')

    args = parser.parse_args()

    file_path = Path(args.file)
    if not file_path.exists():
        print(f"Error: File {args.file} not found", file=sys.stderr)
        sys.exit(1)

    stripper = CommentStripper(keep_docstrings=args.keep_docstrings)

    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()

    stripped = stripper.strip_comments(content, file_path.suffix.lower())

    if args.output:
        with open(args.output, 'w', encoding='utf-8') as f:
            f.write(stripped)
    else:
        print(stripped)

if __name__ == '__main__':
    main()

