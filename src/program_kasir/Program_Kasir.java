/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package program_kasir;
import javax.swing.JOptionPane;
/**
 *
 * @author erik1997
 */
public class Program_Kasir {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // deklarasi
        int belanjaan, diskon, bayar;
        String kartu;
        
        //input
        kartu=JOptionPane.showInputDialog("Apakah Ada Kartu Member?");
        belanjaan=Integer.valueOf(JOptionPane.showInputDialog("Total Belanjaan?"));
        
        //proses
        if (kartu.equalsIgnoreCase("Ya")){
            if (belanjaan > 500000) {
                diskon = 50000;
        } else if (belanjaan > 100000){
            diskon = 15000;
        } else {
            diskon = 0;
        }    
    } else {
        if (belanjaan > 100000) {
            diskon = 5000;
        } else {
            diskon=0;
        }
    }
    
    //total yang harus dibayar
    bayar=belanjaan-diskon;
    JOptionPane.showMessageDialog(null, "Total bayar : Rp. " + bayar);
        
        
    }
}
   