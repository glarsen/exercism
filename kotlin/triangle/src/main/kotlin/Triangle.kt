class Triangle<out T : Number>(a: T, b: T, c: T) {
    val isEquilateral: Boolean = a == b && b == c
    val isIsosceles: Boolean = a == b || a == c || b == c
    val isScalene: Boolean = a != b && b != c

    // Check Triangle Inequality
    init {
        val (x, y, z) = listOf(a.toFloat(), b.toFloat(), c.toFloat())

        val inequality = x > 0.0 && y > 0.0 && z > 0.0 && x + y >= z && x + z >= y && y + z >= x

        require(inequality) { "not a triangle" }
    }
}
