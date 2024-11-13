import sqlite3
import json
import argparse

def sqlite_table_to_json(db_path, table_name, output_file=None):
    """
    Convert a SQLite table to JSON format.

    Args:
        db_path (str): Path to the SQLite database file
        table_name (str): Name of the table to export
        output_file (str, optional): Path to save the JSON file. If None, prints to console.

    Returns:
        list: List of dictionaries containing the table data
    """
    try:
        conn = sqlite3.connect(db_path)
        conn.row_factory = sqlite3.Row
        cursor = conn.cursor()

        cursor.execute(f"PRAGMA table_info({table_name})")
        columns = [column[1] for column in cursor.fetchall()]

        if not columns:
            raise ValueError(f"Table '{table_name}' not found in database")

        cursor.execute(f"SELECT * FROM {table_name}")
        rows = cursor.fetchall()

        result = []
        for row in rows:
            row_dict = dict(row)
            result.append(row_dict)

        if output_file:
            with open(output_file, 'w', encoding='utf-8') as f:
                json.dump(result, f, indent=2, ensure_ascii=False)
            print(f"Data exported to {output_file}")
        else:
            print(json.dumps(result, indent=2, ensure_ascii=False))

        return result

    except sqlite3.Error as e:
        print(f"SQLite error: {e}")
        return None
    except Exception as e:
        print(f"Error: {e}")
        return None
    finally:
        if 'conn' in locals():
            conn.close()

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Convert SQLite table to JSON')
    parser.add_argument('db_path', help='Path to SQLite database file')
    parser.add_argument('table_name', help='Name of the table to export')
    parser.add_argument('-o', '--output', help='Output JSON file path (optional)')

    args = parser.parse_args()
    sqlite_table_to_json(args.db_path, args.table_name, args.output)

