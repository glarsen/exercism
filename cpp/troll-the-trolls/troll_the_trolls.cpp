namespace hellmath {

// TODO: Task 1 - Define an `AccountStatus` enumeration to represent the four
// account types: `troll`, `guest`, `user`, and `mod`.
enum class AccountStatus {
    troll,
    guest,
    user,
    mod
};

// TODO: Task 1 - Define an `Action` enumeration to represent the three
// permission types: `read`, `write`, and `remove`.
enum class Action {
    read,
    write,
    remove
};

// TODO: Task 2 - Implement the `display_post` function, that gets two arguments
// of `AccountStatus` and returns a `bool`. The first argument is the status of
// the poster, the second one is the status of the viewer.
bool display_post(AccountStatus poster, AccountStatus viewer) {
   if (poster == AccountStatus::troll && viewer != AccountStatus::troll) {
       return false;
   }

   return true;
}

// TODO: Task 3 - Implement the `permission_check` function, that takes an
// `Action` as a first argument and an `AccountStatus` to check against. It
// should return a `bool`.
bool permission_check(Action action, AccountStatus status) {
        bool decision;

        switch (status) {
            case AccountStatus::guest:
                if (action == Action::read) {
                    decision = true;
                } else {
                    decision = false;
                }
                break;
            case AccountStatus::troll:
            case AccountStatus::user:
                if (action == Action::remove) {
                    decision = false;
                } else {
                    decision = true;
                }
                break;
            case AccountStatus::mod:
                decision = true;
                break;
        }

        return decision;
}

// TODO: Task 4 - Implement the `valid_player_combination` function that
// checks if two players can join the same game. The function has two parameters
// of type `AccountStatus` and returns a `bool`.
bool valid_player_combination(AccountStatus a, AccountStatus b) {
    // Guests are not allowed
    if (a == AccountStatus::guest || b == AccountStatus::guest) {
        return false;
    }

    // Trolls may only play with other trolls
    if ((a == AccountStatus::troll && b != AccountStatus::troll) || (a != AccountStatus::troll && b == AccountStatus::troll)) {
        return false;
    }

    return true;
}

// TODO: Task 5 - Implement the `has_priority` function that takes two
// `AccountStatus` arguments and returns `true`, if and only if the first
// account has a strictly higher priority than the second.
bool has_priority(AccountStatus first, AccountStatus second) {
        bool decision;

        if (first == second) {
            return false;
        }

        switch (first) {
            case AccountStatus::troll:
                decision = false;
                break;
            case AccountStatus::guest:
                if (second == AccountStatus::troll) {
                    decision = true;
                } else {
                    decision = false;
                }
                break;
            case AccountStatus::user:
            case AccountStatus::mod:
                if (second == AccountStatus::mod) {
                    decision = false;
                } else {
                    decision = true;
                }
                break;
        }

        return decision;
}

}  // namespace hellmath