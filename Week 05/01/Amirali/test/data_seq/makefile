
all: compile1 compile2
	@echo "linking" && echo "done"

compile1:
	@echo "compiling 1"

compile2: compile3
	@echo "compiling 2"

compile3:
	@echo "compiling 3"


all2: clean compile
	@echo "all"
	echo "done"

compile:
	@echo "compile"

clean:
	@echo "clean"

empty:


circle1: compile3 compile2 compile1 compile1
circle2: compile1 compile2 compile3 compile2


invalid_command:
	some_invalid_command_to_cause_failure

invalid_syntax:
	echo "salam'

