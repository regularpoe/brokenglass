require 'dotenv'
require 'haml'
require 'json'
require 'sinatra'

set :haml, :format => :html5

def parse_raw(data)
    ary = data.split("\n")
    { 'artist' => ary[4].gsub('tag artist ', ''), 'album' => ary[5].gsub('tag album ', ''), 'track' => ary[6].gsub('tag title', '') }.to_json
end

get '/' do
    haml :index
end

get '/status' do
    content_type :json
    raw = `cmus-remote --server localhost --passwd foobar -C status`
    parse_raw raw
end

post '/play' do 
    `cmus-remote --server localhost --passwd foobar -p`
    redirect '/'
end

post '/pause' do
    `cmus-remote --server localhost --passwd foobar -u`
    redirect '/'
end

post '/prev' do
    `cmus-remote --server localhost --passwd foobar -r`
    redirect '/'
end

post '/next' do
    `cmus-remote --server localhost --passwd foobar -n`
    redirect '/'
end

post '/stop' do
    `cmus-remote --server localhost --passwd foobar -s`
    redirect '/'
end
