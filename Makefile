.PHONY: viewapi

# 定義要檢查的模組目錄
NODE_MODULES_PATH := node_modules

viewapi:
	@echo "Bootstrapping API document viewer..."
	if [ ! -d "$(NODE_MODULES_PATH)" ]; then \
		echo "Dependencies not installed, installing swagger-ui-express, yamljs..."; \
		[ ! -f "package.json" ] && npm init -y > /dev/null; \
		npm install express swagger-ui-express yamljs; \
	else \
		echo "Dependencies already installed."; \
	fi && \
	echo "Booting up API document viewer..." && \
	node api/view-API-doc.js