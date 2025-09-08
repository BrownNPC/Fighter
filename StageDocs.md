### Documentation for making a stage from within Minecraft


First set FOV to 50, and make sure you have the "Flashback" mod installed.


Then place a barrier block on your left and right so that you can walk 9 blocks from left to right.


Move to the left until you are hugging the left barrier.


Then use the tp command to look stright ahead.

```
tp @p ~ ~ ~ <x angle> 0
```


start the recording from the pause menu.


After 10+ seconds, go to the other side, and look straight ahead, then end the recording.


Now in flashback, you can set up a camera keyframe at the start position and end position.
Make it so it lerps linearly for 200 ticks. ie. from 0s to 1s.

framerate should be 20 fps.

Export mkv, apple prores.

#### Resolution: 384x200 (important)



Use ffmpeg to extract the frames.
```
ffmpeg -i stage.mkv stageFrames/background_%d.png
```

copy the frames into a folder inside assets/stages

run the tool to generate a sheet from the frames.


```
go run tools/bgSprite/bgSprite.go assets/stages/stage1
```


install pngquant and use it to decrease the filesize. for eg.

```
pngquant --quality=95-100 --speed 1 256 --output out-quant.png -- stage_14x15_201frames.png
```

now delete the original stage_NxN_Nframes.png and rename `out-quant.png` to take it's place
