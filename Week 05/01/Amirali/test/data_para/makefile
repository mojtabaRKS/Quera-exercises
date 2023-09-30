
all: compile1 compile2
	@sleep 0.1s
	@echo "linking" && echo "done"

compile1:
	@sleep 0.1s
	@echo "compiling 1"

compile2: compile3
	@sleep 0.1s
	@echo "compiling 2"

compile3:
	@sleep 0.1s
	@echo "compiling 3"


all2: clean compile
	@sleep 0.1s
	@echo "all"
	echo "done"

compile:
	@sleep 0.1s
	@echo "compile"

clean:
	@sleep 0.1s
	@echo "clean"

empty:


circle1: compile3 compile2 compile1 compile1
circle2: compile1 compile2 compile3 compile2


invalid_command:
	some_invalid_command_to_cause_failure


invalid_syntax:
	echo "salam'

