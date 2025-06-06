lint:
	./scripts/lint.sh --sol-lint

fmt-check:
	./scripts/lint.sh --sol-format-check

fmt:
	./scripts/lint.sh --sol-format

clean:
	forge clean

build: clean
	forge build

unit-test: build
	forge test -vvv

analyze:
	slither .

coverage:
	forge coverage --no-match-coverage "script|test|ERC7201"

coverage-debug:
	forge coverage --no-match-coverage "script|test|ERC7201" --report debug

gas-report:
	forge test --gas-report

doc:
	forge doc --serve

generate-bindings: build
	rm -rf abi-bindings
	./scripts/abi_bindings.sh
