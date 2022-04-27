# rand_avatar_gen_go
# A simple go avatar generator.
How it works:
The program uses a preset of images located in the folders 'Base', '1', '2', '3' and so on to generate different avatar combinations using golang's image libraries.

'Base' Folder
- This folder just holds the background JPG.

Numbered Folders
- These folders are basically layers. The first folder, '1', will just hold different variations of the first layer. Different layers will be placed ontop of one another to create a single image in the end.

# To Do
- Make it scalable and more flexible with image sizing, layer counts, more random in the image outputs.
Currently the code is written to work with a specific number of folders (layers) and produces a fixed amount of avatars (currently outputs 10,000 avatars with the set of presets given).
