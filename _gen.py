from collections import OrderedDict

# values copied from the
# https://ru.wikipedia.org/wiki/%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D1%86%D0%B2%D0%B5%D1%82%D0%BE%D0%B2_%D0%B2_X11#%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D1%86%D0%B2%D0%B5%D1%82%D0%BE%D0%B2_[2]
raw = """
Alice Blue 	#F0F8FF 	240 248 255
Antique White 	#FAEBD7 	250 235 215
Aqua 	#00FFFF 	0 255 255
Aquamarine 	#7FFFD4 	127 255 212
Azure 	#F0FFFF 	240 255 255
Beige 	#F5F5DC 	245 245 220
Bisque 	#FFE4C4 	255 228 196
Black 	#000000 	0 0 0
Blanched Almond 	#FFEBCD 	255 235 205
Blue 	#0000FF 	0 0 255
Blue Violet 	#8A2BE2 	138 43 226
Brown 	#A52A2A 	165 42 42
Burlywood 	#DEB887 	222 184 135
Cadet Blue 	#5F9EA0 	95 158 160
Chartreuse 	#7FFF00 	127 255 0
Chocolate 	#D2691E 	210 105 30
Coral 	#FF7F50 	255 127 80
Cornflower 	#6495ED 	100 149 237
Cornsilk 	#FFF8DC 	255 248 220
Crimson 	#DC143C 	220 20 60
Cyan 	#00FFFF 	0 255 255
Dark Blue 	#00008B 	0 0 139
Dark Cyan 	#008B8B 	0 139 139
Dark Goldenrod 	#B8860B 	184 134 11
Dark Gray 	#A9A9A9 	169 169 169
Dark Green 	#006400 	0 100 0
Dark Khaki 	#BDB76B 	189 183 107
Dark Magenta 	#8B008B 	139 0 139
Dark Olive Green 	#556B2F 	85 107 47
Dark Orange 	#FF8C00 	255 140 0
Dark Orchid 	#9932CC 	153 50 204
Dark Red 	#8B0000 	139 0 0
Dark Salmon 	#E9967A 	233 150 122
Dark Sea Green 	#8FBC8F 	143 188 143
Dark Slate Blue 	#483D8B 	72 61 139
Dark Slate Gray 	#2F4F4F 	47 79 79
Dark Turquoise 	#00CED1 	0 206 209
Dark Violet 	#9400D3 	148 0 211
Deep Pink 	#FF1493 	255 20 147
Deep Sky Blue 	#00BFFF 	0 191 255
Dim Gray 	#696969 	105 105 105
Dodger Blue 	#1E90FF 	30 144 255
Firebrick 	#B22222 	178 34 34
Floral White 	#FFFAF0 	255 250 240
Forest Green 	#228B22 	34 139 34
Fuchsia 	#FF00FF 	255 0 255
Gainsboro 	#DCDCDC 	220 220 220
Ghost White 	#F8F8FF 	248 248 255
Gold 	#FFD700 	255 215 0
Goldenrod 	#DAA520 	218 165 32
Gray (X11) 	#BEBEBE 	190 190 190
Gray (W3C) 	#7F7F7F 	127 127 127
Green (X11) 	#00FF00 	0 255 0
Green (W3C) 	#007F00 	0 127 0
Green Yellow 	#ADFF2F 	173 255 47
Honeydew 	#F0FFF0 	240 255 240
Hot Pink 	#FF69B4 	255 105 180
Indian Red 	#CD5C5C 	205 92 92
Indigo 	#4B0082 	75 0 130
Ivory 	#FFFFF0 	255 255 240
Khaki 	#F0E68C 	240 230 140
Lavender 	#E6E6FA 	230 230 250
Lavender Blush 	#FFF0F5 	255 240 245
Lawn Green 	#7CFC00 	124 252 0
Lemon Chiffon 	#FFFACD 	255 250 205
Light Blue 	#ADD8E6 	173 216 230
Light Coral 	#F08080 	240 128 128
Light Cyan 	#E0FFFF 	224 255 255
Light Goldenrod 	#FAFAD2 	250 250 210
Light Gray 	#D3D3D3 	211 211 211
Light Green 	#90EE90 	144 238 144
Light Pink 	#FFB6C1 	255 182 193
Light Salmon 	#FFA07A 	255 160 122
Light Sea Green 	#20B2AA 	32 178 170
Light Sky Blue 	#87CEFA 	135 206 250
Light Slate Gray 	#778899 	119 136 153
Light Steel Blue 	#B0C4DE 	176 196 222
Light Yellow 	#FFFFE0 	255 255 224
Lime 	#00FF00 	0 255 0
Lime Green 	#32CD32 	50 205 50
Linen 	#FAF0E6 	250 240 230
Magenta 	#FF00FF 	255 0 255
Maroon (X11) 	#B03060 	176 48 96
Maroon (W3C) 	#7F0000 	127 0 0
Medium Aquamarine 	#66CDAA 	102 205 170
Medium Blue 	#0000CD 	0 0 205
Medium Orchid 	#BA55D3 	186 85 211
Medium Purple 	#9370DB 	147 112 219
Medium Sea Green 	#3CB371 	60 179 113
Medium Slate Blue 	#7B68EE 	123 104 238
Medium Spring Green 	#00FA9A 	0 250 154
Medium Turquoise 	#48D1CC 	72 209 204
Medium Violet Red 	#C71585 	199 21 133
Midnight Blue 	#191970 	25 25 112
Mint Cream 	#F5FFFA 	245 255 250
Misty Rose 	#FFE4E1 	255 228 225
Moccasin 	#FFE4B5 	255 228 181
Navajo White 	#FFDEAD 	255 222 173
Navy 	#000080 	0 0 128
Old Lace 	#FDF5E6 	253 245 230
Olive 	#808000 	128 128 0
Olive Drab 	#6B8E23 	107 142 35
Orange 	#FFA500 	255 165 0
Orange Red 	#FF4500 	255 69 0
Orchid 	#DA70D6 	218 112 214
Pale Goldenrod 	#EEE8AA 	238 232 170
Pale Green 	#98FB98 	152 251 152
Pale Turquoise 	#AFEEEE 	175 238 238
Pale Violet Red 	#DB7093 	219 112 147
Papaya Whip 	#FFEFD5 	255 239 213
Peach Puff 	#FFDAB9 	255 218 185
Peru 	#CD853F 	205 133 63
Pink 	#FFC0CB 	255 192 203
Plum 	#DDA0DD 	221 160 221
Powder Blue 	#B0E0E6 	176 224 230
Purple (X11) 	#A020F0 	160 32 240
Purple (W3C) 	#7F007F 	127 0 127
Red 	#FF0000 	255 0 0
Rosy Brown 	#BC8F8F 	188 143 143
Royal Blue 	#4169E1 	65 105 225
Saddle Brown 	#8B4513 	139 69 19
Salmon 	#FA8072 	250 128 114
Sandy Brown 	#F4A460 	244 164 96
Sea Green 	#2E8B57 	46 139 87
Seashell 	#FFF5EE 	255 245 238
Sienna 	#A0522D 	160 82 45
Silver 	#C0C0C0 	192 192 192
Sky Blue 	#87CEEB 	135 206 235
Slate Blue 	#6A5ACD 	106 90 205
Slate Gray 	#708090 	112 128 144
Snow 	#FFFAFA 	255 250 250
Spring Green 	#00FF7F 	0 255 127
Steel Blue 	#4682B4 	70 130 180
Tan 	#D2B48C 	210 180 140
Teal 	#008080 	0 128 128
Thistle 	#D8BFD8 	216 191 216
Tomato 	#FF6347 	255 99 71
Turquoise 	#40E0D0 	64 224 208
Violet 	#EE82EE 	238 130 238
Wheat 	#F5DEB3 	245 222 179
White 	#FFFFFF 	255 255 255
White Smoke 	#F5F5F5 	245 245 245
Yellow 	#FFFF00 	255 255 0
Yellow Green 	#9ACD32 	154 205 50
"""

var_names = []
name_vars = OrderedDict()
for line in raw.splitlines():
    line = line.strip()
    if len(line) < 1:
        continue

    name, color_def = line.split('#')

    var_name = name.strip().replace(' ', '', 10).replace('\t', '').replace('(', '').replace(')', '')
    color_name = name.strip()

    var_names.append(var_name)
    name_vars[color_name] = var_name
    print('{0} = X11Color{{"{1}", color.RGBA{{0x{2}, 0x{3}, 0x{4}, 0xFF}}, true}}'.format(
        var_name,
        color_name,
        color_def[0:2],
        color_def[2:4],
        color_def[4:6],
    ))

print('\n\n' + ',\n'.join(var_names) + '\n\n')

for color_name, var_name in name_vars.items():
    print('"{0}": {1},'.format(color_name, var_name))
