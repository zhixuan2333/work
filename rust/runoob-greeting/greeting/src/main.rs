fn main() {
    let list = [50, 100, 150, 200];
    let mut i = 0;
    let index = loop {
        let ch = list[i];
        if ch == 150 {
            break i+1;
        }
        i += 1
    };
    println!("\'150\' index is {}", index)
}
