EXS    :=    ex00 ex01 ex02 ex03 ex04

.PHONY: thru $(EXS)
$(EXS):
	@echo "[[ $@ ]]"
	@mkdir -p $@/vendor/{ft,piscine}
	@cp printrune.go $@/vendor/ft/
	@touch $@/main.go
	@go fmt $@/main.go
	@go fmt $@/vendor/piscine/*
	@cd $@;\
	go mod init $@;\
	go run .

thru: $(EXS)
