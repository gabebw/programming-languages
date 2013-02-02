{--
Load with :l baby.hs:
   $ ghci
   > :l baby.hs
--}
lucky :: (Integral a) => a -> String
lucky 7 = "LUCKY SEVEN"
lucky x = "out of luck :("

addVectors :: (Num a) => (a, a) -> (a, a) -> (a, a)
addVectors (x1, y1) (x2, y2) = (x1 + x2, y1 + y2)

-- fst/snd for triples. _ means "I don't care about it"
first :: (a, b, c) -> a
first (x, _, _) = x

second :: (a, b, c) -> b
second (_, y, _) = y

third :: (a, b, c) -> c
third (_, _, z) = z

head' :: (Show a) => [a] -> a
head' [] = error "Can't call head on an empty list, dummy!"
head' (x:_) = x

-- Check out that x:xs pattern matching
tell :: (Show a) => [a] -> String
tell [] = "The list is empty"
tell (x:[]) = "The list has one element: " ++ show x
tell (x:y:[]) = "The list has two elements: " ++ show x ++ " and " ++ show y
tell (x:y:_) = "This list is long. The first two elements are: " ++ show x ++ " and " ++ show y

length' :: (Num b) => [a] => b
length' [] = 0
lenght' (_:xs) = 1 + length' xs
