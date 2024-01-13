task :run do
  sh "air"
end

def goose(command)
  sh "dotenv -- goose -dir ./db/migration/ #{command}"
end

namespace :db do
  task :gen do
    sh "dotenv -- sqlc generate"
  end

  task :create do
    name = ARGV[1]
    if name.nil? || name.empty? 
      puts "Please provide a name"
      exit 1
    end
    goose "create #{name} sql"

    exit
  end

  task :up do
    goose "up"
  end

  task :status do
    goose "status"
  end

  task :down do
    goose "down"
  end
end

task :default => :run
