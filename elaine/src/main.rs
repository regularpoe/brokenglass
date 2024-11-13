extern crate rand;
use rand::thread_rng;
use rand::Rng;

const QUOTES: &str = "
We don’t know how long this will last. They are a very festive people.
Where’s Kramer?
All right, let’s go, I’ll give you half an hour.
You’re through, Soup Nazi. Pack it up. No more soup for you. Next!
I'm not a lesbian! I hate men, but I'm not a lesbian.
It seems that a psychotic mechanic has absconded with my friend's car.
Well, that's the positive thing about getting sick, you get to lose weight.
Is it possible I’m not attractive as I think I am?
Jerry, it's three-thirty in the morning. I'm at a cockfight. What am I clinging to?
There must be something exciting about this guy if he can arouse that kind of passion.
Yeah, it was, but then I just couldn't decide if he was really sponge-worthy.
Poor little Bubble Boy. He's sitting there waiting for you in his bubble, or igloo thing or whatever.
I think this is the same one I gave him. He recycled this gift. He's a regifter!
You made a man cry? I've never made a man cry. I even kicked a guy in the groin once and he didn't cry.
He took it out.
Maybe the dingo ate your baby.
I don't think George has ever thought he's better than anybody.
Here's to us who wish us well, and those who don't can go to hell.
I had to take a sick day. I'm so sick of these people.";

fn main() {
    let quotes = QUOTES.split("\n").collect::<Vec<&str>>();

    let mut rng = thread_rng();

    let e: usize = rng.gen_range(0..quotes.len());

    println!("{}", quotes[e]);
}
