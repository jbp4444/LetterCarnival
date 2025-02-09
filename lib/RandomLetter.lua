
local globals = require( 'lib.GlobalVars' )

local M = {}

-- relative frequency of all letters
-- -- http://pi.math.cornell.edu/~mec/2003-2004/cryptography/subs/frequencies.html
local all_letters = {"E","T","A","O","I","N","S","R","H","D","L","U","C","M","F","Y","W","G","P","B","V","K","X","Q","J","Z"}
local all_lfreq = {21912,16587,14810,14003,13318,12666,11450,10977,10795,7874,7253,5246,4943,4761,4200,3853,3819,3693,3316,2715,2019,1257,315,205,188,128}
local lfreq_sum = 182303

local letter_lists = {
	{ 'R','F','V', 'U','J','M' },		--   index fingers/single column only
	{ 'T','G','B', 'Y','H','N' },		-- + index fingers/second column
	{ 'E','D','C', 'I','K' },			-- + middle finger/letters only
	{ 'W','S','X', 'O','L' },			-- + ring finger/letters only
	{ 'Q','A','Z', 'P' },				-- + pinkie finger/letters only
	{ '1','2','3','4','5','6','7','8','9','0' }, -- + numbers
}

local last_level = 0
local letters
local letters_cache = {}
local lfreq

function M.shuffle(list)
	for i = 1,#list do
		local j = math.random(#list)
		list[i], list[j] = list[j], list[i]
	end
end

local function calc_lfreq( letters )
	lfreq_sum = 0
	lfreq = { 0 }
	for i = 1,#letters do
		lfreq[i+1] = all_lfreq[ letters[i] ]
		lfreq_sum = lfreq_sum + all_lfreq[ letters[i] ]
	end
	for i = 2,#lfreq do
		lfreq[i] = lfreq[i-1] + lfreq[i]
	end
		for i = 1,#lfreq do
		lfreq[i] = lfreq[i]/lfreq_sum
	end
end
local function select_random_letter()
	local rn = math.random()
	-- print( 'rn=', rn, lfreq[1],lfreq[2] )
	local ltr = 'A'
	for i = 1,#letters do
		if( (rn >= lfreq[i]) and (rn < lfreq[i+1]) ) then
			ltr = letters[i]
			break
		end
	end
	return ltr
end


function M.randomLetter()
	-- prep the underlying structures so we can calc letters
	if( globals.finger_level ~= last_level ) then
		last_level = globals.finger_level
		letters = {}
		for i=1,last_level do
			for j=1,#letter_lists[i] do
				table.insert( letters, letter_lists[i][j] )
			end
		end
		--calc_lfreq( letters )
		local txt = ''
		for i=1,#letters do
			txt = txt .. letters[i]
		end
		print( 'letters for level:'..last_level..'='..txt )
	end

	if( #letters_cache == 0 ) then
		-- make 5 copies of the current letter-set
		letters_cache = {}
		for i=1,3 do
			for j=1,#letters do
				table.insert( letters_cache, letters[j] )
			end
		end
		M.shuffle( letters_cache )
	end

	ltr = table.remove( letters_cache )
	return ltr
end

return M
