#include <opencv2/opencv.hpp>

using namespace cv;
using namespace std;

// Image matrices
Mat src, dst;

// Blob detector parameters
SimpleBlobDetector::Params params;

int main(int argc, char **argv) {
    src = imread(argv[1], IMREAD_GRAYSCALE);

    params.minThreshold = 10;
	params.maxThreshold = 200;

	// Filter by Area.
	params.filterByArea = true;
	params.minArea = 5;

	// Filter by Circularity
	params.filterByCircularity = true;
	params.minCircularity = 0.1;

	// Filter by Convexity
	params.filterByConvexity = true;
	params.minConvexity = 0.87;

	// Filter by Inertia
	params.filterByInertia = true;
	params.minInertiaRatio = 0.01;

    // Coordinate array
    vector<KeyPoint> keypoints;

    SimpleBlobDetector::create(params)
        ->detect(src, keypoints);

    drawKeypoints(src, keypoints, dst, Scalar(0, 0, 255), DrawMatchesFlags::DRAW_RICH_KEYPOINTS);

    imwrite("out.jpg", dst);

    return 0;
}