<?php
    require './vendor/autoload.php';
    class Redis {
        private $client;

        public function __construct() {
            Predis\Autoloader::register();

            $this->client = new Predis\Client(array(
                    "scheme" => "tcp",
                    "host" => "127.0.0.1",
                    "port" => 6379,
                    "password" => "evenkingscry!")) or die;
            
        }
        

        public function addRequestToQueue($request) {
            $currentTime = time();
            $this->client->set($currentTime, $request);
        }
    }
?>