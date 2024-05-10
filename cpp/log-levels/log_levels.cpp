#include <string>

namespace log_line {
    const std::string LevelError = "ERROR";
    const std::string LevelInfo = "INFO";
    const std::string LevelWarning = "WARNING";

    std::string message(const std::string& line) {
        return line.substr(line.find(':') + 2);
    }

    std::string log_level(const std::string& line) {
        return line.substr(1, line.find(']') - 1);
    }

    std::string reformat(const std::string& line) {
        return message(line) + " (" + log_level(line) + ")";
    }
}
