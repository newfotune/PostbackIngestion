<?php
   require_once("RequestObject.php");
   require("RedisInstance.php");

    $response = array();
    $revieved = json_decode(file_get_contents('php://input'), true);

    if (!isset($revieved["endpoint"], $revieved["data"]))  {
        http_response_code(400);
        return;
    }
    
    $request = new Request($revieved["endpoint"], $revieved["data"]);
    
    $redis = new Redis();
    if ($redis->addRequestToQueue($request->getRequestJSON())) {
        $response['error'] = false;
        $response['message'] = "Successfully queued request";
    } else {
        $response['error'] = true;
        $response['message'] = "Failed to add request to Redis";
    }
    echo json_encode($response);
?>