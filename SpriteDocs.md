#### Docs for rendering sprites using Block Bench


#### Import minecraft character into blockbench:

- New Minecraft skin
- Select the model you want
- File > convert > generic model
- then save


#### Animation tips
- Select all keyframes, and set lerp mode to `smooth`
- rotate the character 45 degrees on the Y axis.

#### How to export:
- Right click viewport > Save camera angle
- Right click viewport > angles > <your saved angle> > edit
- Copy the following settings:
---
- Projection: Orthographic
- Camera Position: 512 22 0
- Focal Point: 0 22 0
- Zoom: 0.28


#### Then use that camera angle.

- View > Record GIF
- Copy these settings:
---
- Format: PNG Sequence
- Length Mode: Animation Length
- FPS: (use an odd number eg. 11 or 23)
- Resolution: 512 512
- Check "Play Animation"
- Confirm


#### Converting rendered sprite to sprite sheet:
- Extract the `.zip` you exported from Block Bench.
- Run the tool like so:


```
go run tools/bgSprite/bgSprite.go -downscale -mode sprite -prefix <YourAnimationName eg. walk> -o ./assets/characters/steve path/to/frames
```
