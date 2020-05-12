dockerbuild_email_service:
	cd email_service && docker build -t abjnet_email_service . 

dockerbuild_email_service_proxy:
	cd email_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_email_service . 

dockerbuild_user_service:
	cd user_service && docker build -t abjnet_user_service . 

dockerbuild_user_service_proxy:
	cd user_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/ --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_user_service . 

dockerbuild_product_service:
	cd product_service && docker build -t abjnet_product_service . 

dockerbuild_product_service_proxy:
	cd product_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_product_service . 

dockerbuild_souscription_service:
	cd souscription_service && docker build -t abjnet_souscription_service . 

dockerbuild_souscription_service_proxy:
	cd souscription_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_souscription_service . 

dockerbuild_taksrunner_service:
	cd taksrunner_service && docker build -t abjnet_taksrunner_service . 

dockerbuild_taksrunner_service_proxy:
	cd taksrunner_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_taksrunner_service . 