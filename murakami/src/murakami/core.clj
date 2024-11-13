(ns murakami.core
  (:gen-class)
  (:require [markdown.core :as markdown]
            [selmer.parser :as selmer]
            [clojure.java.io :as io]
            [clojure.string :as str]))

(def template "
<!DOCTYPE html>
<html lang=\"en\">
<head>
  <meta charset=\"utf-8\">
  <title></title>
  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">
  <link rel=\"stylesheet\" href=\"../css/normalize.css\">
  <link rel=\"stylesheet\" href=\"../css/skeleton.css\">
  <link rel=\"stylesheet\" href=\"../css/style.css\">
  <link rel=\"stylesheet\" href=\"../css/github.min.css\">
  <script src=\"../highlight.min.js\"></script>
  <link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"/apple-touch-icon.png\">
  <link rel=\"icon\" type=\"image/png\" sizes=\"32x32\" href=\"/favicon-32x32.png\">
  <link rel=\"icon\" type=\"image/png\" sizes=\"16x16\" href=\"/favicon-16x16.png\">
  <link rel=\"manifest\" href=\"/site.webmanifest\">
</head>
<body>
    <div class=\"container\">
        <div class=\"row\">
            <header><h1>a title goes here</h1></header>
        </div>
        <div class=\"row post\">
            {{content | safe}}
        </div>
    </div>
    <script>hljs.highlightAll();</script>
</body>
</html>
")

(defn is-markdown-file? [file]
  (str/ends-with? (.getName (io/file file)) ".md"))

(defn markdown-to-html-file [file]
  (let [content (slurp file)
        html-content (markdown/md-to-html-string content)
        rendered-html (selmer/render template {:content html-content})
        output-file (str (str/replace (str (.getName file)) #"\.md$" ".html"))]
    (spit output-file rendered-html)
    (println "Written to" output-file)))

(defn read-markdown-files [dir]
  (let [files (filter is-markdown-file? (file-seq (io/file dir)))]
    (doseq [file files]
      (markdown-to-html-file file))))

(defn -main [& args]
  (if-let [dir (first args)]
    (read-markdown-files dir)
    (println "Provide path to Markdown files")))
