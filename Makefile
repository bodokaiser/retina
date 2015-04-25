INCLUDE := -I /usr/local/include/opencv
LIBRARY := -L /usr/local/lib \
	-l opencv_core \
	-l opencv_imgproc \
	-l opencv_highgui \
	-l opencv_imgcodecs \
	-l opencv_features2d

build:
	$(CXX) $(INCLUDE) $(LIBRARY) src/main.cpp -o main.out	

clean:
	rm main.out out.jpg