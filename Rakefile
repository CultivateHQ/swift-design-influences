require 'rake/clean'
require 'rouge'
require 'erb'
require 'rack'
require 'webrick'

PROJECT_ROOT = File.dirname(__FILE__)
SRC_DIR = File.join(PROJECT_ROOT, 'src')
WEB_DIR = File.join(PROJECT_ROOT, 'public')
BUILD_DIR = File.join(PROJECT_ROOT, 'build')

directory BUILD_DIR

LANG_SRC = FileList["#{SRC_DIR}/*"]
WEB_SRC = FileList["#{WEB_DIR}/js/*.js", "#{WEB_DIR}/css/*.css"]
BUILT_FILES = FileList["#{BUILD_DIR}/index.html"]

def highlight(infilename, syntax)
  infile = File.read(infilename)
  Rouge.highlight(infile, syntax, 'html')
end

def build_index_html
  swift_example = highlight("#{SRC_DIR}/Swift.swift", 'swift')
  scala_example = highlight("#{SRC_DIR}/Scala.scala", 'scala')
  go_example = highlight("#{SRC_DIR}/Go.go", 'go')

  index_erb = ERB.new(File.read("#{WEB_DIR}/index.html.erb"))
  index_erb.result(binding)
end

WEB_SRC.each do |file_path|
  built_path = file_path.pathmap("%{^#{WEB_DIR},#{BUILD_DIR}}p")

  BUILT_FILES.include(built_path)

  file built_path => [BUILD_DIR, file_path] do
    mkdir_p File.dirname(built_path)
    cp file_path, built_path
  end
end

file "#{BUILD_DIR}/index.html" => [*LANG_SRC, BUILD_DIR] do |t|
  index_html = build_index_html
  File::write(t.name, index_html)
end

desc "Run a server instance for development convenience"
task :server do |t|
  builder = Rack::Builder.new do
    run -> (env) { [200, {"Content-Type" => "text/html"}, [build_index_html] ] }
    use Rack::Static, :urls => ["/css", "/js"], :root => WEB_DIR
  end

  trap 'INT' do Rack::Handler::WEBrick.shutdown end
  Rack::Handler::WEBrick.run builder, :Port => 3000
end

desc "Process and output the project to the #{BUILD_DIR} folder"
task :build => BUILT_FILES do
  puts "Built."
end

desc "Runs build"
task :default => :build

desc "Print out project paths and other info"
task :info do
  puts "PROJECT_ROOT: #{PROJECT_ROOT}"
  puts "SRC_DIR: #{SRC_DIR}"
  puts "WEB_DIR: #{WEB_DIR}"
  puts "BUILD_DIR: #{BUILD_DIR}"
end

CLEAN.include(BUILT_FILES)
CLOBBER.include(BUILD_DIR)
