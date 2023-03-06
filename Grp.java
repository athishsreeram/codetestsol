import java.util.*;

public class Grp {

    private Map<Integer,Map<Integer,Integer>> graph;
    private boolean[] visited;

    public Grp(int size){
        graph =  new HashMap<>();
        visited = new boolean[size];
        for(int i=0;i<=size;i++){
            graph.put(i,new HashMap());
        }
    }

    public void addNode(int u,int v,int weight,boolean directed){
        if(directed) {
            Map adjNode = graph.get(u);
            adjNode.put(v, weight);
        }else{
            Map adjNode = graph.get(v);
            adjNode.put(u, weight);
        }
    }

    public void DFS(Map<Integer,Map<Integer,Integer>> graph,int src){
        visited[src-1]=true;
        System.out.println(" "+src);
        Map adjNode = graph.get(src);
        if(adjNode != null) {
            adjNode.forEach( (v,w) -> {
                if(!visited[(Integer)v-1]) {
                    DFS(graph, (Integer) v);
                }
            });
        }
    }

    public void BFS(Map<Integer,Map<Integer,Integer>> graph,int source){

        boolean[] visited = new boolean[graph.size()];

        visited[source-1]=true;


        LinkedList<Integer> queue = new LinkedList();
        queue.add(source);

        while (queue.size() != 0) {

            Integer src = queue.poll();

            System.out.println(" "+src);

            Map adjNode = graph.get(src);
            if (adjNode != null) {
                adjNode.forEach((v, w) -> {
                    if (!visited[(Integer) v - 1]) {
                       queue.add((Integer) v);
                    }
                });
            }

        }
    }

    public static void main(String[] args) {

        Grp grp = new Grp(4);
        grp.addNode(1,2,1,true);
        grp.addNode(1,3,1,true);
        grp.addNode(2,4,1,true);

        grp.DFS(grp.graph,1);

        System.out.println(" ");

        grp.BFS(grp.graph,1);
    }



}
