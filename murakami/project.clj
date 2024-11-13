(defproject murakami "1.0.0"
  :description "Convert MD to HTML template"
  :url "https://github.com/regularpoe/murakami"
  :license {:name "GNU GENERAL PUBLIC LICENSE Version 3"
            :url ""}
  :dependencies [[org.clojure/clojure "1.11.1"]
                 [org.clojure/tools.namespace "1.5.0"]
                 [markdown-clj "1.10.0"]
                 [selmer "1.12.61"]]
  :main ^:skip-aot murakami.core
  :target-path "target/%s"
  :profiles {:uberjar {:aot :all
                       :jvm-opts ["-Dclojure.compiler.direct-linking=true"]}})
