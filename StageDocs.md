### Documentation for making a stage from within Minecraft

#### Flashback mod is required.

### To record:

- Stage width: 5 blocks (use barriers to contain the width)
- Must be symmetric.
- use the tp command to look stright ahead.

```
tp @p ~ ~ ~ <x angle> 0
```

- Start flashback mod recording from pause menu.
- Hug the left most barrier.
- After 5s hug the right most barrier
- Save the recording
- Save and quit > click the camera icon in the main menu.
- Right click the player, hide during export
- Make sure the player is hugging the left most barrier
- And then spectate the player > Add a camera keyframe
- Spectate player when theyre hugging the right most barrier > spectate > add a camera keyframe



#### Export as video from flashback mod.
- Fov:60
- recording must be 59 ticks
- framerate 20 (so that 20*3s = 60)
- start tick 0, end tick 59
- export png sequence with the format `background_%d` in flashback mod.
- **Resolution: 1536x1200(important)**


Anti-alias the frames and convert to sprite sheet
```
go run tools/bgSprite/bgSprite.go -downscale  -mode bg -o ./assets/stages/stage1 path/to/frames
```

install pngquant and use it to decrease the filesize. for eg.
```
pngquant --quality=95-100 --speed 1 256 --output out-quant.png -- stage_8x8_60frames.png
```

now delete the original stage_NxN_Nframes.png and rename `out-quant.png` to take it's place
