<?php

// Turn on strict mode here.
error_reporting(E_STRICT);

class KentikAlert { 


    public $type;            // assign variable for the alert type (Alarm - alert has been triggered / Clear - alert is no longer active)
    public $severity;        // assign variable for the severity as defined for the alert in the portal - Info, Major, Minor, Critical    
    public $alert_id;        // assign variable for this specific alert notification 
    public $event_id;
    public $key_name;        // assign variable for key_name and value as defined in the SQL query
    public $key_value;
    public $device;          // Device name.       
    public $sup1_outport;    // assign variable for Layer 4 destination port for supplemental one (two returned values)
    public $sup2_outport;
    public $dst_addr;        // assign variable for IP destination address for only supplemental two (two returned values);
    public $sup1_dst;
    public $sup2_dst;
    public $sup1_inint;
    public $sup2_inint;
    
    public function __construct($raw) {
        
        $parsed = json_decode($raw, true);
        if ($parsed == NULL) {
            throw new Exception('Invalid input: ' . $raw);            
        }
        
        // Now, pick alert apart.
        $this->type = $parsed["type"]; 
        $this->severity = $parsed["severity"];
        $this->alert_id = $parsed["alert_id"];
        $this->event_id = $parsed["event_id"];
        $this->key_name = $parsed["key_name"];
        $this->key_value = $parsed["key_value"];        
        $this->device = $parsed['query_result']['i_device_name'];
        $this->sup1_outport = $parsed['supl_sql_one_value']['0']['l4_dst_port'];
        $this->sup2_outport = $parsed['supl_sql_one_value']['1']['l4_dst_port'];
        $this->sup1_dst = $parsed['supl_sql_two_value']['0']['ipv4_dst_addr'];
        $this->sup2_dst = $parsed['supl_sql_two_value']['1']['ipv4_dst_addr'];
        $this->sup1_inint = $parsed['supl_sql_two_value']['0']['input_port'];
        $this->sup2_inint = $parsed['supl_sql_two_value']['1']['input_port'];
    }

    // Magic function, called when evaluating this class as a string.
    public function __toString ( ) {

        return implode(PHP_EOL, array(
                                      $this->type,
                                      $this->severity,                                      
                                      "alert_id: " . $this->alert_id,
                                      "event_id: " . $this->event_id,
                                      "key name: " . $this->key_name,
                                      "key value: " . $this->key_value,
                                      "device_name: " . $this->device,
                                      $this->sup1_outport,
                                      $this->sup2_outport,
                                      $this->sup1_inint,
                                      $this->sup2_inint,
                                      $this->sup1_dst,
                                      $this->sup2_dst,
                                      PHP_EOL
                                      )
                       );        
    }
}

// Output file to write data to
$myFile = "/tmp/testFile.txt";

// Listen for HTTP post and extract/decode JSON body
$aRequest = "";

// If command line, read from stdin.
if (php_sapi_name() == "cli") {
    $aRequest = new KentikAlert(trim(file_get_contents('php://stdin')));
} else {
    // Otherwise, read the body of the POST.
    $aRequest = new KentikAlert(file_get_contents(trim('php://input')));
}

//Print everything to a file
file_put_contents($myFile,$aRequest, FILE_APPEND | LOCK_EX);

// Respond with a success response.
echo '{ "success": true }' . PHP_EOL;

?>
