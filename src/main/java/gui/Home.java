package gui;

import javax.swing.*;

public class Home extends JFrame {

    private JPanel panel1;
    private JButton button1;

    public Home(String title){
        this.setVisible(true);
        this.setContentPane(new Home("Loopbox").panel1);
        this.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        this.pack();
    }

    public static void main(String[] args) {
        System.out.println("Starting Loopbox GUI.");
        new Home("Loopbox");
    }
}
