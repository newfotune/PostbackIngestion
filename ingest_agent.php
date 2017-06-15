<?php
    //phpinfo();

   require './vendor/autoload.php';

   Predis\Autoloader::register();

   $client = new Predis\client();
   $client->set('foo', 'bar');
   $value = $client->get('foo');

   echo $value
?>