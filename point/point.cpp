#include <cmath>
#include <memory>
#include <iostream>

#include "point.hpp"
#include "utils.h"
#include "_cgo_export.h"

Point::Point(double x, double y)
{
    Log(1, string_format("new point (%p) x=%2.f y=%2.f", this, x, y));

    this->_x = x;
    this->_y = y;
}

Point::~Point() {
    // std::cout << this << std::endl;
};

double Point::x()
{
    return this->_x;
}

double Point::y()
{
    return this->_y;
}

double Point::distance(const std::unique_ptr<Point> & p, const std::unique_ptr<Point> & q)
{
    double dx = p->x() - q->x();
    double dy = p->y() - q->y();

    return sqrt(dx * dx + dy * dy);
}
