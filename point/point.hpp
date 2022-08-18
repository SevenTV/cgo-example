#ifndef POINT_HPP
#define POINT_HPP

#include <memory>

class Point {
public:
    Point(double x, double y);
    ~Point();
    double x();
    double y();
    static double distance(const std::unique_ptr<Point>& p, const std::unique_ptr<Point>& q);

private:
    double _x;
    double _y;
};

#endif /* POINT_HPP */
