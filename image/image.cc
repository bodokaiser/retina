#include "image.h"

#include <opencv2/opencv.hpp>

using namespace cv;
using namespace std;

Mat src, dst;

SimpleBlobDetector::Params params;

Mat image_to_mat(image* img) {
    vector<uchar> buf(img->bytes, img->bytes + img->length);

    return imdecode(buf, IMREAD_GRAYSCALE);
}

image* mat_to_image(Mat mat) {
    vector<uchar> buf;
    imencode(".jpeg", mat, buf);

    uchar* bytes = (uchar*) malloc(buf.size());
    copy(buf.begin(), buf.end(), bytes);

    image* img = (image*) malloc(sizeof(image));
    img->bytes = bytes;
    img->length = buf.size();
    img->format = (char *)"jpeg";

    return img;
}

image* process(image* img) {
    src = image_to_mat(img);

    params.minThreshold = 0.4;

    params.filterByArea = true;
    params.minArea = 14;
    params.maxArea = 800;

    params.filterByConvexity = true;
    params.minConvexity = 0.8;

    vector<KeyPoint> keypoints;

    SimpleBlobDetector::create(params)
        ->detect(src, keypoints);

    drawKeypoints(src, keypoints, dst, Scalar(0, 255, 0), DrawMatchesFlags::DRAW_RICH_KEYPOINTS);

    return mat_to_image(dst);
}