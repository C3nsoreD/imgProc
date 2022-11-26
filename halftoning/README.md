# Halftoning 

Here we explore what halftoning is and learn about some image procesing. Its a nice technique were images are reduced to dot of 
varying size giving an artistic look. Initially it was used in priniting to reduce ink usage.

The process:
 - first step is to take a jpg image and convert that to grayscale.
 - That grayscale image contains gray colors ranging from 0(black) -> 255(white)
 - Compare the gray image with a threshold (127) and returning white or black. lower than returns black.

 The image returned in step 3 is basically quantized containing "very contrasted".