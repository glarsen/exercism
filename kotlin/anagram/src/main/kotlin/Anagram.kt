class Anagram(private val source: String) {
    private fun analyze(word: String): Map<Char, List<Char>> =
            word.lowercase().groupBy { it }

    private val match: Map<Char, List<Char>> = analyze(source)

    fun match(anagrams: Collection<String>): Set<String> =
        anagrams.filter {
             candidate ->
                source.lowercase() != candidate.lowercase() && match == analyze(candidate)
        }.toSet()
}
