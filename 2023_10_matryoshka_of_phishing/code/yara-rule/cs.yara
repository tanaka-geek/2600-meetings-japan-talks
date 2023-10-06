rule cobaltstrike_beacon_strings
{
meta:
    author = "Elastic"
    description = "Identifies strings used in Cobalt Strike Beacon DLL."
strings:
    $a = "%02d/%02d/%02d %02d:%02d:%02d"
    $b = "Started service %s on %s"
    $c = "%s as %s\\%s: %d"
condition:
    1 of them
}