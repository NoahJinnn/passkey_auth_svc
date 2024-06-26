#!/opt/homebrew/bin/bash
brew install bash
shopt -s globstar
exec 2>/dev/null
cd ..
sloc=$(cat ./**/*.go | wc -l)
sloc_gen=$(cat ./**/*.*.go | wc -l)
sloc_gen_api=$(cat ./api/**/*.*.go | wc -l)
sloc_gen_mock=$(cat ./**/mock.*.go | wc -l)
sloc_gen_other=$((sloc_gen - sloc_gen_api - sloc_gen_mock))
sloc_gen_test=$(cat ./**/*.*_test.go | wc -l)
sloc_test=$(($(cat ./**/*_test.go | wc -l) - sloc_gen_test))
sloc_code=$((sloc - sloc_test - sloc_gen))
printf "SLOC: all=%d (code=%d test=%d generated all=%d (api=%d mock=%d other=%d))\n" \
	"$sloc" "$sloc_code" "$sloc_test" \
	"$sloc_gen" "$sloc_gen_api" "$sloc_gen_mock" "$sloc_gen_other"
