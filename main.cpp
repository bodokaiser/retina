#include <opencv2/opencv.hpp>

using namespace cv;
using namespace std;

// Image matrices
Mat src, dst;

// Blob detector parameters
SimpleBlobDetector::Params params;

int main(int argc, char **argv) {
    src = imread(argv[1], IMREAD_GRAYSCALE);

    params.minThreshold = 0.4;

    params.filterByArea = true;
    params.minArea = 14;
    params.maxArea = 800;

    params.filterByConvexity = true;
    params.minConvexity = 0.8;

    // Coordinate array
    vector<KeyPoint> keypoints;

    SimpleBlobDetector::create(params)
        ->detect(src, keypoints);

    drawKeypoints(src, keypoints, dst, Scalar(0, 255, 0), DrawMatchesFlags::DRAW_RICH_KEYPOINTS);

    imwrite(argv[2], dst);

    return 0;
}