#ifndef UTILS_HPP
#define UTILS_HPP

#include <memory>
#include <stdexcept>
#include <string>

// this include is exported by cgo and it provides the Log function
#include "_cgo_export.h"

template <typename... Args>
char* string_format(const std::string& format, Args... args)
{
    int size = std::snprintf(nullptr, 0, format.c_str(), args...) + 1; // Extra space for '\0'
    if (size <= 0)
        throw std::runtime_error("Error during formatting.");

    auto buf = new char[size];

    std::snprintf(buf, size, format.c_str(), args...);

    return buf;
}

template <typename... Args>
void GoLog(const int level, const std::string& format, Args... args) {
    auto msg = string_format(format, args...);

    // also note, that this function here is expensive to call back into golang from c
    // you should try avoid doing this logging is a good usecase if u have a global golang logger, but try be very selective with where u log.
    // since we create so many points logging on a new point each time costs a lot of time.
    // this is just a debug tool so we can see that it actually works.
    Log(level, msg);

    free(msg);
}

#endif /* UTILS_HPP */
