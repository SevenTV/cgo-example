#include "cpoint.h"

#include "point.hpp"
#include <cassert>
#include <cstdlib>
#include <memory>

struct point_t {
    Point obj;
};

point_t* point_new(double x, double y)
{
    point_t* pt = new point_t{Point(x, y)};

    return pt;
}

void point_delete(point_t* self)
{
    assert(self);

    // if you used any pointers in the struct you must clean them up here, since we called new in the point_new func we must delete the pointer we created.
    // since this is a c api we cannot use unique pointers unfortunately go does not interact directly with them.
    // this is okay however because in go we can simply add finalizers which call this delete function for us, doing the memory management on our behalf. 

    delete self;
}

double point_x(point_t* self)
{
    assert(self);

    return self->obj.x();
}

double point_y(point_t* self)
{
    assert(self);

    return self->obj.y();
}

double point_distance(point_t* p, point_t* q)
{
    assert(p);
    assert(q);

    return Point::distance(p->obj, q->obj);
}
