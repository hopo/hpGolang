<?php

class MyFileObject{
    function isFile(){
        return is_File($this->$filename);
    }
}

$file = new MyFileObject();
$file->filename = 'data.txt';
var_dump($file->isFile());
var_dump($file->filename);

$file2 = new MyFileObject();
$file2->filename = 'data2.txt';
var_dump($file2->isFile());
var_dump($file2->filename);

?>
