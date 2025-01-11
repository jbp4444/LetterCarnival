-- base functions for all targets

local myrandom = require( 'lib.RandomLetter' )

-- can use socket.gettime() too:
math.randomseed( os.time() )

local targetlib = {}

function targetlib.new( in_args )
	local obj = {}
	obj.params = {
		position = {0,0,0},
		scale = {1,1,1},
		speed = 150,
		fg_images = {'default'},
		bg_image = 'default',
	}

	-- override with input params
	if( in_args ~= nil ) then
		for k,v in pairs(obj.params) do
			obj.params[k] = in_args[k] or obj.params[k]
		end
	end

	-- state variables for this object
	obj.flipped = false
	obj.active = true

	-- set up some gfx-obj params
	local scl = vmath.vector3( obj.params.scale[1],obj.params.scale[2],obj.params.scale[3] )
	go.set_scale( scl )
	local pos = vmath.vector3( obj.params.position[1],obj.params.position[2],obj.params.position[3] )
	go.set_position( pos )

	-- set up list of percentages for each fg image
	local fg_pcts = {0}
	for i = 2,#obj.params.fg_images do
		fg_pcts[i] = fg_pcts[i-1] + 1/(#obj.params.fg_images)
	end
	table.insert( fg_pcts, 1 )
	obj.fg_pcts = fg_pcts

	local rn = math.random()
	local img
	for i = 1,#obj.fg_pcts do
		if( (rn >= obj.fg_pcts[i]) and (rn < obj.fg_pcts[i+1]) ) then
			img = obj.params.fg_images[i]
			break
		end
	end
	obj.image = img
	
	rn = math.random()
	local stick_img
	if( rn < 0.333 ) then
		stick_img = 'stick_wood_outline'
	elseif( rn < 0.666 ) then
		stick_img = 'stick_woodFixed_outline'
	else
		stick_img = 'stick_metal_outline'
	end
	obj.stick_image = stick_img

	msg.post( '#front', 'play_animation', {id=hash(img)} )
	msg.post( '#back', 'play_animation', {id=hash(obj.params.bg_image)} )
	msg.post( '#stick', 'play_animation', {id=hash(stick_img)} )
	msg.post( '#back', 'disable' )
	--msg.post( '#xmark', 'disable' )

	-- get a random letter based on what level was selected
	ltr = myrandom.randomLetter()
	obj.letter = ltr
	local clr = vmath.vector4( 0,0,0,1 )
	label.set_text( '#label', ltr )
	go.set( '#label', 'color', clr )
-- 	local scl = vmath.vector3( 2,2,2 )
-- 	go.set( '#label', 'scale', scl )

	msg.post( '.', 'acquire_input_focus' )

	print( 'targetlib:'..tostring(img) )

	--  --  --  --  --  --  --  --  --  --
	  --  --  --  --  --  --  --  --  --
	--  --  --  --  --  --  --  --  --  --

	function obj:done_flip_anim()
		print( 'done_flip_anim' )
		go.delete()
	end
	function obj:mid_flip_anim()
		print( 'mid_flip_anim '..tostring(obj) )
		msg.post( '#back', 'enable' )
		msg.post( '#front', 'disable' )
		msg.post( '#circle', 'disable' )
		msg.post( '#label', 'disable' )
		go.set( '.', 'scale.x', 0 )
		go.animate( '.', 'scale.x', go.PLAYBACK_ONCE_FORWARD, -1.0*obj.params.scale[1], go.EASING_LINEAR, 0.1, 0, self.done_flip_anim )
	end

	function obj:update( dt )
		local pos = go.get_position()
		pos.x = pos.x + dt*self.params.speed
		go.set_position( pos )

		if( self.active and ((pos.x > 1111) or (pos.x < -111)) ) then
			print( 'target complete' )
			self.active = false
			local pos = go.get_position()
			msg.post( '/gamelogic', 'complete', {hit=self.flipped, pos=pos, sender_id=go.get_id()} )
		end
	end

	function obj:on_input( action_id, action )
		-- Add input-handling code here
		local rtn = false
		if( self.active and (action_id ~= nil) ) then
			if( action.text ~= nil ) then
				if( self.flipped == false ) then
					local ltr = string.upper( action.text )
					if( ltr == self.letter ) then
						print( 'starting flip_anim' )
						sound.play( '#sound' )
						msg.post( '/gamelogic', 'score', {score=5,pos=go.get_position()} )
						go.animate( '.', 'scale.x', go.PLAYBACK_ONCE_FORWARD, 0.0, go.EASING_LINEAR, 0.1, 0, self.mid_flip_anim )
						self.flipped = true
						rtn = true
					end
				end
			end
		end
		return rtn
	end

	return obj
end

return targetlib
