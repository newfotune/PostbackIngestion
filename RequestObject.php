<?php

    class Request {
        private $request;

        public function __construct($theEndPoint, $theData) {
            $this->request['method'] = $theEndPoint["method"];
            $this->request['url'] = $theEndPoint["url"];
            $this->request['data'] = $theData;
        }

        public function getRequestJSON() {
            return json_encode($this->request);
        }
    }
?>
