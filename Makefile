build_sensor_image:
	docker build ./internship-task-go-master/cmd/sensor -t youla_dev_internship_task_go_sensor:latest

build_service_image:
	docker-compose -f ./sensor-checker/docker-compose.yml build

run_service: build_sensor_image build_service_image
	docker-compose -f ./sensor-checker/docker-compose.yml up
