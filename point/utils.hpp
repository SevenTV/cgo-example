#ifndef UTILS_HPP
#define UTILS_HPP

#include <memory>
#include <stdexcept>
#include <string>

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

#endif /* UTILS_HPP */
