const express = require('express');
const app = express();
const swaggerUi = require('swagger-ui-express');
const YAML = require('yamljs');
const path = require('path');

// 設定你的 yaml 檔案路徑
// 假設這個 js 檔在根目錄，而 yaml 在 api 資料夾內
const filePath = path.join(__dirname, 'openapi.yml');

try {
    // 嘗試讀取檔案
    const swaggerDocument = YAML.load(filePath);
    
    // 設定路由
    app.use('/docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument));
    
    console.log('==================================================');
    console.log(`文件伺服器已啟動！`);
    console.log(`請用瀏覽器打開: http://localhost:3000/docs`);
    console.log('==================================================');

} catch (e) {
    console.error("讀取 yaml 檔案失敗，請檢查路徑是否正確：", filePath);
    console.error(e);
}

app.listen(3000);