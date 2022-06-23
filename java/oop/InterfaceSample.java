package oop;

public class InterfaceSample {
    public static void main(String[] args) {
        System.out.println("light");
        CarLight carLight = new CarLight();
        carLight.turnOn();
        carLight.turnOff();
    }
    
}
interface Light {
    void turnOn();

    void turnOff();
}

class CarLight implements Light {

    @Override
    public void turnOn() {
        System.out.println("turn on");
    }

    @Override
    public void turnOff() {
        System.out.println("turn off");
    }
}


