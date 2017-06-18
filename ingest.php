<?php
    //phpinfo();
   require_once("RequestObject.php");
   require("RedisInstance.php");

   /*Predis\Autoloader::register();

   $client = new Predis\client();
   $client->set('foo', 'bar');
   $value = $client->get('foo');

   echo $value*/
    $revieved = json_decode(file_get_contents('php://input'), true);

    if (!isset($revieved["endpoint"], $revieved["data"]))  {
        http_response_code(400);
        return;
    }
    
    $request = new Request($revieved["endpoint"], $revieved["data"]);
    //echo $request->getRequestJSON();
    
    $redis = new Redis();
    $redis->addRequestToQueue($request->getRequestJSON());
?>