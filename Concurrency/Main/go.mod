module example.com/main

go 1.20

//replace example.com/routine => ../Routine

replace example.com/chanels => ../Chanels

require example.com/chanels v0.0.0-00010101000000-000000000000
