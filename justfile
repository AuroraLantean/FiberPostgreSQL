#!/usr/bin/env just --justfile
# https://github.com/casey/just

set dotenv-required
set dotenv-load
set quiet
#set export   ... to export all variables below as the environment variables
alias b:= build
alias t:= test
dirpath:= "/mnt/sda3/"
py := "python3.13"
backup_dir := "./aarchives"
filenamePrefix := "a"
time1 := datetime("%Y_%m_%d-%H_%M_%S")
filename := filenamePrefix+time1+".tar.gz"

default:
	echo "hello time: {{time1}}"
	just --list
init:
  go mod init newModuleName
setup:
  cp $GIT_LOCAL_HOOKS_PATH .git/hooks/
  chmod a+x .git/hooks//pre-commit
cleancache: 
	go clean -i -x -modcache
remove:
	rm go.mod go.sum
fmt:
	go fmt
tidy:
  go mod tidy
build:tidy
  go build -o main *.go
install:
  go install
add:
  go get github.com/charmbracelet/log@latest
update:
	go get -u
	go mod tidy
term:
  terminator -l '12GoFiber1' > /dev/null 2>&1 &  
run:
	go run *.go
# to run and watch file changes... depending on  watchexec: cargo binstall watchexec-cli
runm selection:
  @echo 'Run {{selection}} ...'
  go run *.go
watch:
	watchexec just run
run2:
	task -w dev
test:
	go test
slumber:
	RUST_LOG=slumber=debug slumber
makesolwallet:
  solana-keygen new --outfile ./wallets/keypair1.json
download:
	go mod download
get1:
	curl -i localhost:3000/books/show?isbn=978-1505255607
post1:
	curl -X POST -H 'Content-Type: application/json' -d "{\"test\": \"that\"}"

deploy: 
	anchor deploy
bk:
  echo "bk to run"
  echo "backup_dir = {{backup_dir}}"
  echo "filename = {{filename}}"
  echo "filenamePrefix = {{filenamePrefix}}"
  echo "time1 = {{time1}}"

  echo "before executing compression command..."
  tar cpzf {{backup_dir}}/{{filename}} .env* .gitignore justfile LICENSE *.*
  #CODEOWNERS build .prettier*

  echo "backup completed successfully."
  ls {{backup_dir}}/{{filenamePrefix}}*
checkcommits:
  git log
env:
	source .env
js:
  #!/usr/bin/env node
  console.log('Greetings from JavaScript!')

_private:
	echo "this can only be invoked by other receipe"
	-failed_but_continue_next
	echo "skipping the error"
test_error: _private
	echo "run test_error"

more:
  firefox https://github.com/casey/just
# https://just.systems/man/en/