#include "cpoint.h"

#include "point.hpp"
#include <cassert>
#include <cstdlib>
#include <memory>

struct point_t {
    std::unique_ptr<Point> obj;
};

point_t* point_new(double x, double y)
{
    point_t* pt = new point_t {
        std::make_unique<Point>(x, y)
    };

    return pt;
}

void point_delete(point_t* self)
{
    assert(self);

    delete self;
}

double point_x(point_t* self)
{
    assert(self);

    return self->obj->x();
}

double point_y(point_t* self)
{
    assert(self);

    return self->obj->y();
}

double point_distance(point_t* p, point_t* q)
{
    assert(p);
    assert(q);

    return Point::distance(p->obj, q->obj);
}
