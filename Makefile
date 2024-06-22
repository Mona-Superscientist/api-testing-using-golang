.PHONY: build up exec test allure-report clean

build:
	docker-compose build

up:
	docker-compose up

bash:
	docker-compose exec app /bin/bash

test:
	make run-test

run-test:
	ginkgo -r --keep-going --junit-report=report.xml -v

allure-report:
	mkdir allure-results && mv report.xml allure-results && allure generate

allure-open:
	allure open

clean:
	rm -rf ./allure-results ./allure-report
