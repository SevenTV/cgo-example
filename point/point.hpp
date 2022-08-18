#ifndef POINT_HPP
#define POINT_HPP

#include <memory>

class Point {
public:
    Point(double x, double y);
    ~Point();
    double x() const;
    double y() const;
    static double distance(const Point& p, const Point& q);

private:
    double _x;
    double _y;
};

#endif /* POINT_HPP */
