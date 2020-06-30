package gui;

import javax.swing.*;

public class Home extends JFrame {

    private JPanel homePanel;
    private JButton button1;

    public Home(String title) {
        super(title);
        this.pack();

        this.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        this.setVisible(true);
        this.setSize(100,100);
    }
}
