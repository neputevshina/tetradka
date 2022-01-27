#!/usr/bin/fish

function newest -a path
	if test -f $path
		stat $path -c %Y
		return
	end
	
	set -x max 0		# you can't have a js file dated jan 1 1970
	for i in $path/*	# accurate enough
		if test (stat $i -c %Y) -gt $max 
			set max (stat $i -c %Y)
		end
	end
	echo $max
end

function older -a a b
	test (newest $a) -lt (newest $b)
end

if older build.fish app
	fish -c 'cd app; vue build'
	# Replace default title
	set temp /tmp/(date +%s).html~
	cp app/dist/index.html $temp
	sed 's/Vue CLI App/Tetradka/g' <$temp >app/dist/index.html
	rm $temp
end
go build
touch build.fish
