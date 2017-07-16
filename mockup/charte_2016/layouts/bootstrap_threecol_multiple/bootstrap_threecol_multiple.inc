<?php

/**
 * Implements hook_panels_layouts().
 */
// Plugin definition
$plugin = array(
  'title' => t('Three Columns Multiple'),
  'category' => t('Bootstrap'),
  'icon' => 'bootstrap_threecol_multiple.png',
  'theme' => 'bootstrap_threecol_multiple',
  'admin css' => 'bootstrap_threecol_multiple.css',
  'regions' => array(
    'top' => t('Top'),
    'left1' => t('Left1'),
    'middle1' => t('Middle1'),
    'right1' => t('Right1'),
    'left2' => t('Left2'),
    'middle2' => t('Middle2'),
    'right2' => t('Right2'),
    'left3' => t('Left3'),
    'middle3' => t('Middle3'),
    'right3' => t('Right3'),
    'bottom' => t('Bottom'),
  ),
);
