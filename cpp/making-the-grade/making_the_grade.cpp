#include <array>
#include <string>
#include <vector>

// Round down all provided student scores.
std::vector<int> round_down_scores(std::vector<double> student_scores) {
    return {student_scores.begin(), student_scores.end()};
}


// Count the number of failing students out of the group provided.
int count_failed_students(std::vector<int> student_scores) {
    int count = 0;

    for (auto score: student_scores) {
        if (score <= 40) {
            count++;
        }
    }

    return count;
}

// Determine how many of the provided student scores were 'the best' based on the provided threshold.
std::vector<int> above_threshold(const std::vector<int>& student_scores, int threshold) {
    std::vector<int> best;

    for (auto score: student_scores) {
        if (score >= threshold) {
            best.emplace_back(score);
        }
    }

    return best;
}

// Create a list of grade thresholds based on the provided highest grade.
std::array<int, 4> letter_grades(int highest_score) {
    int interval = (highest_score - 40) / 4;

    std::array<int, 4> thresholds = {
            41,
            41 + interval,
            41 + interval * 2,
            41 + interval * 3,
            };

    return thresholds;
}

// Organize the student's rank, name, and grade information in ascending order.
std::vector<std::string> student_ranking(std::vector<int> student_scores, std::vector<std::string> student_names) {
    int count = student_scores.size();

    std::vector<std::string> rankings;
    rankings.reserve(count);

    for (int i=0; i < count; i++) {
        rankings.emplace_back(std::to_string(i + 1) + ". " + student_names[i] + ": " + std::to_string(student_scores[i]));
    }

    return rankings;
}

// Create a string that contains the name of the first student to make a perfect score on the exam.
std::string perfect_score(std::vector<int> student_scores, std::vector<std::string> student_names) {
    for(int i=0; i < student_scores.size(); i++){
        if (student_scores[i] == 100) {
            return student_names[i];
        }
    }

    return "";
}
