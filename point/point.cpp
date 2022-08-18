#include <cmath>
#include <iostream>
#include <memory>

#include "point.hpp"
#include "utils.hpp"

Point::Point(double x, double y)
{
    GoLog(1, "new point (%p) x=%2.f y=%2.f", this, x, y);

    this->_x = x;
    this->_y = y;
}

Point::~Point() = default;

double Point::x() const
{
    return this->_x;
}

double Point::y() const
{
    return this->_y;
}

double Point::distance(const Point& p, const Point& q)
{
    double dx = p.x() - q.x();
    double dy = p.y() - q.y();

    return sqrt(dx * dx + dy * dy);
}
